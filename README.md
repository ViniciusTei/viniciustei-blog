# Blog

A simple web application that serves articles written in Markdown format. The application will read the Markdown files, convert them into HTML, and serve them as web pages.

The goal is to create the application without any extern dependecys, using only the Go standard library and the Browser's built-in capabilities.

### Features:

- [x] Read articles from Markdown files.
- [ ] Convert Markdown into HTML.

  - [x] Parse Header to h1, h2, h3, etc.
  - [x] Parse paragraphs.
  - [x] Parse lists.
  - [x] Parse text formatting (bold, italic, etc.).
  - [x] Parse links.
  - [x] Parse images.
  - [x] Parse code blocks.
  - [ ] Parse blockquotes.
  - [ ] Parse tables.

- [x] Serve articles as web pages.
- [x] List all articles on a home page.

### Project Structure:

- main.go: Entry point of the application.
- articles/: Directory containing Markdown files.
- templates/: Directory containing HTML templates.
- handlers/: Directory for HTTP handlers (business logic).
- tests/: Directory for test files.
