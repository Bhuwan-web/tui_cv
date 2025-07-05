package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle()

type component struct {
	list_items []list.Item
	title      string
}

// defaultStruct should return an empty component for placeholder, not initialComponent.
func defaultStruct() component {
	return component{}
}

type item struct {
	name           string
	desc           string
	nested_component func() component
	prev_component   func() component
}

// SetItem is a generic setter for convenience if needed, but not strictly used here.
func SetItem(name, desc string, nested_component, prev_component func() component) item {
	return item{name: name, desc: desc, nested_component: nested_component, prev_component: prev_component}
}

// SetRootItem defines an item that leads to a nested menu.
// Its prev_component should be the component that *contains* this root item.
// For top-level root items (like those in initialComponent), prev_component can be a function returning an empty component or nil.
func SetRootItem(name, desc string, nested_component func() component) item {
	return item{name: name, desc: desc, nested_component: nested_component, prev_component: func() component { return component{} }} // Placeholder
}

// SetLeafItem defines an item that just displays information, with no further nesting.
// Its prev_component must be explicitly set to the component it belongs to.
func SetLeafItem(name, desc string, prev_component func() component) item {
	return item{name: name, desc: desc, prev_component: prev_component, nested_component: func() component { return component{} }}
}

func (i item) Title() string       { return i.name }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.name }

type model struct {
	list        list.Model
	currentPath []component // Stack to keep track of the navigation path
	large_desc  bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.large_desc = false
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch value := msg.String(); value {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			selectedItem := m.list.SelectedItem()
			if selectedItem == nil {
				return m, nil // No item selected, do nothing
			}

			if i, ok := selectedItem.(item); ok {
				if i.nested_component().title != "" { // Check if it has a valid nested component
					// Navigate into the nested component
					nextComponent := i.nested_component()
					m.currentPath = append(m.currentPath, m.listToComponent()) // Push current component to path
					m.list.SetItems(nextComponent.list_items)
					m.list.Title=nextComponent.title
					m.list.Select(0) // Select the first item in the new list
					return m, nil
				}else if len(i.desc)>m.list.Width()-4 { // Check if description is too long
					m.large_desc = true
					return m,nil
				}
			}

		case "esc":
			if len(m.currentPath) > 0 {
				// Pop the last component from the path to go back
				prevComponent := m.currentPath[len(m.currentPath)-1]
				m.currentPath = m.currentPath[:len(m.currentPath)-1] // Remove from stack

				m.list.SetItems(prevComponent.list_items)
				m.list.Title = prevComponent.title
				m.list.Select(0) // Select the first item in the previous list
				return m, nil
			} else {
				// If currentPath is empty, we are at the top level, so quit
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) LongDescView() string {
	selectedItem := m.list.SelectedItem().(item)
	titleStyle := lipgloss.NewStyle().Bold(true).Underline(true)
	descStyle := lipgloss.NewStyle().Width(m.list.Width()-4)
	title := titleStyle.Render(selectedItem.Title())
	desc := descStyle.Render(selectedItem.Description())
	return docStyle.Render(title + "\n\n" + desc + "\n\n↑/k up • ↓/j down • q quit")
}
func (m model) View() string {
	if m.large_desc {
		return m.LongDescView()
	}
	return docStyle.Render(m.list.View())
}

// Helper to convert the current list state back into a component
func (m model) listToComponent() component {
	return component{
		list_items: m.list.Items(),
		title:      m.list.Title,
	}
}

// --- CV Components ---

func initialComponent() component {
	return component{
		list_items: []list.Item{
			SetRootItem("Introduction", "Basic Information and Contact", IntroComponent),
			SetRootItem("Skills", "Technical and Non-Technical Proficiencies", SkillsComponent),
			SetRootItem("Experience", "Details of Professional Experience", ExperienceComponent),
			SetRootItem("Projects", "Details of Self-Employed Projects", ProjectsComponent),
			SetRootItem("Education", "Academic Background", EducationComponent),
			SetRootItem("Languages", "Language Proficiencies", LanguagesComponent),
		},
		title: "Bhuwan Panta's CV",
	}
}

func IntroComponent() component {
	// Reference to the current initialComponent for prev_component
	initial := initialComponent // Capture the function reference

	return component{
		list_items: []list.Item{
			SetLeafItem("Contact", "Email: ricky.pantha@gmail.com | Phone: +977-9844718578 | Location: Lalitpur, Nepal", initial),
			SetLeafItem("Summary", "Results-driven and solution-oriented Software Engineer adept at analyzing and developing software to achieve scalable solutions. Skilled in fostering collaboration, optimizing services, and delivering projects on time and within scope.", initial),
			SetLeafItem("LinkedIn", "linkedin.com/in/bhuwan-panta", initial),
			SetLeafItem("Github", "github.com/bhuwan-panta", initial),
		},
		title: "Introduction",
	}
}

func SkillsComponent() component {
	initial := initialComponent
	return component{
		list_items: []list.Item{
			SetLeafItem("Programming Languages", "Python, Go, JavaScript", initial),
			SetLeafItem("Frameworks", "FastAPI, Flask, Django, Node (Runtime), Express", initial),
			SetLeafItem("Databases", "MongoDB, MySQL, PostgreSQL", initial),
			SetLeafItem("Cloud Services", "AWS (S3, Lambda)", initial),
			SetLeafItem("Search Technologies", "Vector Search, Elasticsearch, Fuzzy Search", initial),
			SetLeafItem("Other (Technical)", "REST API, OPA, Rego, SOAP XML, Docker, Kubernetes, RabbitMQ, Microservices", initial),
			SetLeafItem("Other (Non-Technical)", "Client Communication, Product Initiatives, Leadership", initial),
		},
		title: "Skills",
	}
}

func ExperienceComponent() component {
	// When setting SetRootItem for Experience, its nested component's prev_component
	// should point back to this ExperienceComponent itself.
	expComp := ExperienceComponent // Capture the function reference
	return component{
		list_items: []list.Item{
			SetRootItem("Software Engineer, Tekvortex", "Bhaktapur, Nepal | Jan 2021 – Jan 2022", func() component { return TekvortexExperienceComponent(expComp) }),
			SetRootItem("Software Engineer, RippeyAI", "Louisville, CO | Aug 2022 – Jul 2024", func() component { return RippeyAIExperienceComponent(expComp) }),
			SetRootItem("Software Engineer, LancemeUp", "Lalitpur, Nepal | Feb 2022 – Jul 2022", func() component { return LancemeUpExperienceComponent(expComp) }),
		},
		title: "Experience",
	}
}

// Pass the *function* of the parent component
func RippeyAIExperienceComponent(parentComponent func() component) component {
	return component{
		list_items: []list.Item{
			SetLeafItem("Deployment Free API Integration", "Implemented Domain-Driven Design; onboarded 10 customers/8 carriers; consolidated microservices (20% cost reduction); reduced onboarding time by 60%.", parentComponent),
			SetLeafItem("Microsoft Teams Integration", "Constructed seamless MS Teams integration for chatbots; reduced customer onboarding to 30 mins; implemented real-time failover (50% response time reduction).", parentComponent),
			SetLeafItem("Enhanced Email Parsing Service", "Improved email parsing for foreign characters (5% data integrity); increased efficiency for diverse attachments (15%); increased overall accuracy by 20%.", parentComponent),
			SetLeafItem("Developed Internal Tools", "Enhanced accuracy for date/currency formats (10% error reduction); created Universal Unit Conversion tool (~35% processing speed increase); developed vector search for charge codes (30% search accuracy); automated Excel ops (80% time reduction).", parentComponent),
			SetLeafItem("Built Customizable Rate Engine", "Analyzed libraries (15% dependency cost reduction); engineered Rate Engine (40% config time reduction); collaborated with CTO (20% timeline reduction); participated in meetings (30% project success rate).", parentComponent),
		},
		title: "RippeyAI - Software Engineer",
	}
}

func LancemeUpExperienceComponent(parentComponent func() component) component {
	return component{
		list_items: []list.Item{
			SetLeafItem("Enhanced Legacy Project", "Maintained backend for Medisoft independently; understood codebase quickly (50% transition time reduction); interacted with clients (25% client satisfaction increase).", parentComponent),
		},
		title: "LancemeUp - Software Engineer",
	}
}
func TekvortexExperienceComponent(parentComponent func() component) component {
	// When setting SetRootItem for Experience, its nested component's prev_component
	// should point back to this ExperienceComponent itself.
	return component{
		list_items: []list.Item{
			SetRootItem("MITM Proxy", "Working on BAF(Build Application Firewall) Product", parentComponent),
		},
		title: "Tekvortex - Software Engineer",
	}
}

func ProjectsComponent() component {
	projComp := ProjectsComponent // Capture the function reference
	return component{
		list_items: []list.Item{
			SetRootItem("Share Excel", "Nawalparasi, Nepal | Feb 2020 - Apr 2021", func() component { return ShareExcelProjectComponent(projComp) }),
			SetRootItem("Automatic Grade Ledger Application", "Nawalparasi, Nepal | Feb 2020 - Apr 2021", func() component { return GradeLedgerProjectComponent(projComp) }),
		},
		title: "Solopreneur Projects",
	}
}

func ShareExcelProjectComponent(parentComponent func() component) component {
	return component{
		list_items: []list.Item{
			SetLeafItem("Description", "Developed stock portfolio management using Excel/VBA for ~100 users.", parentComponent),
			SetLeafItem("Features", "Live portfolio tracking, watchlist, interactive dashboards, near real-time stock prices (40% user engagement).", parentComponent),
			SetLeafItem("Methodology", "Applied Agile; continuously improved based on feedback (1000% product value increase).", parentComponent),
		},
		title: "Share Excel Project",
	}
}

func GradeLedgerProjectComponent(parentComponent func() component) component {
	return component{
		list_items: []list.Item{
			SetLeafItem("Description", "Developed an application to assist teachers in publishing results from home during COVID (benefited over 10 teachers).", parentComponent),
			SetLeafItem("Extended Scope", "Built comprehensive School Management Application (admission, accounting, result management); decreased application usage by 60% (likely meant increased efficiency/reduced manual work).", parentComponent),
		},
		title: "Automatic Grade Ledger Application",
	}
}

func EducationComponent() component {
	initial := initialComponent
	return component{
		list_items: []list.Item{
			SetLeafItem("Bachelor of Computer Science and Information Technology", "Bhaktapur Multiple Campus, Bhaktapur, Nepal | Apr 2021 – Apr 2025", initial),
			SetLeafItem("Science / Physics", "Tilottama Higher Secondary School, Butwal, Nepal | Aug 2018 – Sept 2020", initial),
		},
		title: "Education",
	}
}

func LanguagesComponent() component {
	initial := initialComponent
	return component{
		list_items: []list.Item{
			SetLeafItem("Nepali", "Native", initial),
			SetLeafItem("English", "Advanced", initial),
			SetLeafItem("Hindi", "Conversational", initial),
			SetLeafItem("German", "Basic", initial),
		},
		title: "Languages",
	}
}

func main() {
	initialComp := initialComponent()
	m := model{
		list:        list.New(initialComp.list_items, list.NewDefaultDelegate(), 0, 0),
		currentPath: []component{}, // Initialize an empty path stack
	}
	m.list.Title=initialComp.title
	m.list.Styles.TitleBar.Align(lipgloss.Center)
	m.list.SetFilteringEnabled(false)
	m.list.SetShowStatusBar(false)

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		
		panic(err)
	}
}