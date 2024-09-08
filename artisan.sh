#!/bin/bash

# Define colors
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to check if Go is installed
check_go_installed() {
    if command -v go >/dev/null 2>&1; then
        echo -e "${BLUE}Go is installed.${NC}"
        return 0
    else
        echo -e "${RED}Go is not installed. Please install Go first.${NC}"
        return 1
    fi
}

# Function to check if a package binary is in PATH
check_package_installed() {
    local package="$1"
    if command -v "$package" >/dev/null 2>&1; then
        echo -e "${BLUE}Package $package is installed.${NC}"
        return 0
    else
        echo -e "${RED}Package $package is not installed.${NC}"
        return 1
    fi
}

# Function to install packages from a file
install_packages() {
    local file="$1"
    while IFS= read -r line; do
        if [[ $line =~ ^go\ install ]]; then
            local package=$(echo "$line" | awk '{print $2}')
            echo -e "${BLUE}Installing package $package...${NC}"
            go install "$package"
            if [ $? -eq 0 ]; then
                echo -e "${BLUE}Package $package installed successfully.${NC}"
            else
                echo -e "${RED}Failed to install package $package.${NC}"
            fi
        fi
    done < "$file"
}

# Function to uninstall packages from a file
uninstall_packages() {
    local file="$1"
    while IFS= read -r line; do
        if [[ $line =~ ^go\ install ]]; then
            local package=$(echo "$line" | awk '{print $2}')
            local binary_name=$(basename "$package")
            echo -e "${BLUE}Uninstalling package $package...${NC}"
            rm -f "$HOME/go/bin/$binary_name"
            if [ $? -eq 0 ]; then
                echo -e "${BLUE}Package $package uninstalled successfully.${NC}"
            else
                echo -e "${RED}Failed to uninstall package $package.${NC}"
            fi
        fi
    done < "$file"
}

# Main script logic
if [ "$1" == "install" ]; then
    if [ "$2" == "-r" ]; then
        if [ -f "$3" ]; then
            check_go_installed
            if [ $? -eq 0 ]; then
                install_packages "$3"
                # Check if the package is installed and in PATH
                while IFS= read -r line; do
                    if [[ $line =~ ^go\ install ]]; then
                        local package_name=$(echo "$line" | awk '{print $2}' | sed 's/github.com\/[^\/]*\/\([^@]*\)/\1/')
                        # Extract the binary name from the package
                        local binary_name=$(basename "$package_name")
                        check_package_installed "$binary_name"
                    fi
                done < "$3"
            fi
        else
            echo -e "${RED}File $3 does not exist.${NC}"
        fi
    else
        echo -e "${RED}Invalid option. Use -r to specify the file.${NC}"
    fi
elif [ "$1" == "uninstall" ]; then
    if [ "$2" == "-r" ]; then
        if [ -f "$3" ]; then
            uninstall_packages "$3"
        else
            echo -e "${RED}File $3 does not exist.${NC}"
        fi
    else
        echo -e "${RED}Invalid option. Use -r to specify the file.${NC}"
    fi
else
    echo -e "${RED}Invalid command. Use install or uninstall.${NC}"
fi
