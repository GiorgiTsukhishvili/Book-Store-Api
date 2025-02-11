# BookShelf

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Introduction

BookShelf is a platform designed for businesses to showcase their books along with prices. Users can browse through the collection, add books to their favorites, and leave reviews. This project aims to provide a comprehensive solution for book management and user engagement.

## Features

- Businesses can add new books with prices
- Users can browse and search for books
- Users can add books to their favorites
- Users can leave reviews for books
- Organize books by genres

## Installation

To install and run BookShelf locally, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/bookshelf.git
   ```
2. Navigate to the project directory:
   ```sh
   cd bookshelf
   ```
3. Install the dependencies:
   ```sh
   go mod tidy
   ```
4. Start the application:
   ```sh
   go run main.go
   ```

## Testing

To test application locally, run follow these steps:

1. Start the application:
   ```sh
   go run main.go
   ```
2. Run tests from test folder:
   ```sh
   go test ./tests/... -v
   ```

## Usage

Once the application is running, you can access it at `http://localhost:3000`. Use the interface to add, remove, and organize your books.

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Open a pull request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

If you have any questions or suggestions, feel free to contact me on [LinkedIn](https://www.linkedin.com/in/giorgi-tsukhishvili-05b213201/).
