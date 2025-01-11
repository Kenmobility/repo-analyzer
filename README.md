# GitHub Repository Analyzer

This Go application analyzes the structure and file sizes of a public GitHub repository.

## Features
- Clones a public GitHub repository.
- Recursively analyzes the file structure.
- Outputs the structure and sizes in JSON format.

## Usage
1. Clone the project:
   ```bash
   git clone https://github.com/yourusername/github-analyzer.git

2. Build the application:
   ```bash
   go build -o analyzer

3. Run the application:
   ```bash
   ./analyzer <repository-url>

## Tests
   Run tests with 
   ```bash
   go test ./... -v
   ```


## Example Output
 ```json
    {
        "clone_url": "https://github.com/exampls/example.git",
        "size": "155.23 KB",
        "files": [
            {
                "name": ".env.example",
                "size": "232 B"
            },
            {
                "name": ".gitignore",
                "size": "42 B"
            }
        ],
        "folders": [
            {
                "name": "src",
                "files": [
                    {
                        "name": "main.go",
                        "size": "1.23 KB"
                    }
                ],
                "folders": [],
            }
        ],
    }
 ```


