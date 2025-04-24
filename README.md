# Blog

### Features:
        
- [x] Read articles from Markdown files.
- [ ] Convert Markdown into HTML.
    - [x] Parse Header to h1, h2, h3, etc.
    - [x] Parse paragraphs.
    - [x] Parse lists.
    - [ ] Parse text formatting (bold, italic, etc.).
    - [ ] Parse links.
    - [ ] Parse code blocks.
    - [ ] Parse blockquotes.
    - [ ] Parse tables.
    - [ ] Parse images.
- [x] Serve articles as web pages.
- [x] List all articles on a home page.

### Project Structure:

- main.go: Entry point of the application.
- articles/: Directory containing Markdown files.
- templates/: Directory containing HTML templates.
- handlers/: Directory for HTTP handlers (business logic).
- tests/: Directory for test files.
