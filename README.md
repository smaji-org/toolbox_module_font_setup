# Font Setup

`font_setup` is a utility module designed to install and uninstall fonts from the system across different platforms (Unix-like and Windows). It ensures that fonts are correctly copied to the system's font directories and registered with the operating system's font management systems.

## Features

- **Cross-platform Support**: Works on Linux, Unix-like systems and Windows.
- **Automatic Registration**: 
  - On **Unix-like systems**, it automatically triggers `fc-cache` to update the font cache.
  - On **Windows**, it uses a dedicated tool (`cjkv_toolbox_win_font.exe`) to register each font file.
- **Clean Uninstallation**: Provides an uninstaller that removes the installed fonts and cleans up the registration.

## Project Structure

- `installer/`: Contains the logic and executable for installing fonts.
- `uninstaller/`: Contains the logic and executable for removing fonts.
- `data/`: (Expected) Directory containing the `.otf` and `.ttf` font files to be installed.
- `Makefile`: Build instructions for different environments.
- `package.sh`: Script to package the project into a compressed archive.

## Prerequisites

- **Go**: Required for building the tools.
- **Windows Specific**: Requires `cjkv_toolbox_win_font.exe` to be present in the toolbox directory for font registration.

## Getting Started

### Build

To build the tools for the current platform and Windows in debug mode:
```bash
make debug
```

To build the release versions (optimized and with specific flags):
```bash
make release
```

### Packaging

To package the project with font data:
```bash
make package DATA=./path/to/your/font/data
```

### Distribution

The primary goal of this project is to package the `installer`, `uninstaller`, and the font data into individual components for distribution. 

After running the `make package` command, the `package` directory will contain the final artifacts organized by Operating System and Architecture:

`package/<OS>/<ARCH>/font_setup.tgz`

These `.tgz` files are the distribution components. Each one contains the necessary executables and font data tailored for a specific platform.

### Using as a Template

This project can serve as a template for specific font management modules. If you want to customize your own font management module:

1. Set the `project` variable in `package.sh` to your desired project name.
2. Prepare your font files and use the directory containing them as the `DATA` directory during packaging.

After running `make package DATA=./path/to/your/font/data`, you will find the font management components for each OS/Architecture in the `package/` directory with the following structure:
`package/<OS>/<ARCH>/<project_name>.tgz`

## Platform Specifics

### Unix-like Systems
- Fonts are installed to `~/.local/share/fonts`.
- `fc-cache` is called automatically to refresh the system font cache.

### Windows Systems
- Fonts are installed to `%LocalAppData%\Microsoft\Windows\Fonts`.
- Each font is registered using `cjkv_toolbox_win_font.exe`.

## License

See the `LICENSE` file for more information.
