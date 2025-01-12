# GitHub Repository Analyzer

This Go application analyzes the structure and file sizes of a public GitHub repository.

## Features
- Clones a public GitHub repository.
- Recursively analyzes the file structure.
- Outputs the structure and sizes in JSON format.

## Usage
1. Clone the project:
   ```bash
   git clone https://github.com/kenmobility/repo-analyzer.git

2. CD into the project:
   ```bash
   cd repo-analyzer

3. Build the application:
   ```bash
   go build -o analyze

4. Run the application with any public github repository-url as argument:
   ```bash
   ./analyze <repository-url>

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


