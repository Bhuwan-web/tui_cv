# TUI CV Portfolio

A Terminal User Interface (TUI) based interactive CV/portfolio application built with Go, showcasing professional experience, skills, and projects in a navigable menu system.

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Terminal](https://img.shields.io/badge/Terminal-%23054020?style=for-the-badge&logo=gnu-bash&logoColor=white)

## 🚀 Features

- **Interactive Navigation**: Navigate through different sections using arrow keys
- **Hierarchical Menu System**: Drill down into detailed information with nested menus
- **Responsive Design**: Automatically adjusts to terminal window size
- **Long Description View**: Automatically switches to expanded view for lengthy descriptions
- **Breadcrumb Navigation**: Easy navigation back to previous sections using ESC key
- **Cross-Platform**: Builds for Windows, Linux, and macOS

## 📋 Sections

The CV is organized into the following main sections:

- **Introduction**: Contact information, summary, and social links
- **Skills**: Technical and non-technical proficiencies
- **Experience**: Detailed professional work history
- **Projects**: Self-employed and personal projects
- **Education**: Academic background
- **Languages**: Language proficiencies

## 🛠️ Technologies Used

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)**: TUI framework for Go
- **[Bubbles](https://github.com/charmbracelet/bubbles)**: Common TUI components
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)**: Style definitions for terminal UIs
- **Go**: Programming language

## 🏗️ Architecture

The application follows a component-based architecture:

### Core Components

- **`model`**: Main application state containing the list model and navigation stack
- **`component`**: Represents a menu section with items and title
- **`item`**: Individual menu items with navigation capabilities

### Item Types

- **Root Items**: Lead to nested menus (e.g., "Experience" → individual job details)
- **Leaf Items**: Display information without further nesting

### Navigation System

- Uses a stack-based approach (`currentPath`) to track navigation history
- Supports forward navigation (Enter) and backward navigation (ESC)
- Maintains context and state across navigation levels

## 🚦 Getting Started

### Prerequisites

- Go 1.19 or higher
- Terminal with Unicode support for best experience

### Installation

1. Clone the repository:
```bash
git clone https://github.com/bhuwan-panta/tui-cv-portfolio.git
cd tui-cv-portfolio
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build and run:
```bash
make run
```

## 📦 Build Options

The project includes a Makefile with multiple build targets:

### Local Development
```bash
# Build and run locally
make run

# Build only
make build
```

### Cross-Platform Builds
```bash
# Build for Windows
make build_win

# Build for Linux
make build_linux
```

### Manual Build Commands
```bash
# Default build
go build -o cv main.go

# Windows build
GOOS=windows GOARCH=amd64 go build -o cv.exe main.go

# Linux build
GOOS=linux GOARCH=amd64 go build -o cv_linux main.go
```

## 🎮 Usage

### Navigation Controls

| Key | Action |
|-----|--------|
| `↑/k` | Move up in the menu |
| `↓/j` | Move down in the menu |
| `Enter` | Select item / Navigate into submenu |
| `ESC` | Go back to previous menu |
| `Ctrl+C` | Exit application |
| `q` | Quit (when in long description view) |

### Navigation Flow

1. **Start**: Main CV menu with primary sections
2. **Select**: Use arrow keys to highlight desired section
3. **Enter**: Press Enter to navigate into the selected section
4. **Explore**: Browse through detailed information
5. **Return**: Use ESC to go back to previous menu level
6. **Exit**: Use Ctrl+C to quit the application

## 🎨 Customization

### Adding New Sections

1. Create a new component function following the pattern:
```go
func NewSectionComponent() component {
    return component{
        list_items: []list.Item{
            SetLeafItem("Item Name", "Item Description", initialComponent),
            // Add more items...
        },
        title: "Section Title",
    }
}
```

2. Add the section to `initialComponent()`:
```go
SetRootItem("New Section", "Description", NewSectionComponent),
```

### Styling

Modify the `docStyle` and create new styles using Lipgloss:
```go
var customStyle = lipgloss.NewStyle().
    Foreground(lipgloss.Color("205")).
    Background(lipgloss.Color("235")).
    Padding(1, 2)
```

## 🏃‍♂️ Performance Features

- **Lazy Loading**: Components are generated on-demand
- **Memory Efficient**: Uses function references instead of storing all data
- **Responsive**: Automatically adapts to terminal size changes
- **Fast Navigation**: Stack-based navigation for instant back/forward movement

## 🐛 Error Handling

The application includes robust error handling for:
- Invalid navigation states
- Empty menu selections
- Terminal size changes
- Application lifecycle management

## 📝 Code Structure

```
├── main.go              # Main application entry point
├── Makefile            # Build automation
├── go.mod              # Go module dependencies
└── README.md           # This file
```

### Key Functions

- `initialComponent()`: Main menu definition
- `*Component()`: Individual section components
- `SetRootItem()`: Creates navigable menu items
- `SetLeafItem()`: Creates information display items
- `Update()`: Handles user input and state changes
- `View()`: Renders the current UI state

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 About

Created by **Bhuwan Panta** as a learning project for TUI development in Go.

- **Email**: ricky.pantha@gmail.com
- **LinkedIn**: [linkedin.com/in/bhuwan-panta](https://linkedin.com/in/bhuwan-panta)
- **GitHub**: [github.com/bhuwan-panta](https://github.com/bhuwan-panta)

## 🙏 Acknowledgments

- [Charm](https://charm.sh/) for the amazing TUI libraries
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) community for inspiration
- Go community for excellent tooling and resources

---

*Built with ❤️ using Go and Bubble Tea*