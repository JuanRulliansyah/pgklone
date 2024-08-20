# pgklone

**pgklone** is a straightforward tool for cloning PostgreSQL databases. With **pgklone**, you can easily duplicate a PostgreSQL database (Database A) to another (Database B) by simply providing the source database URL. This tool is designed for ease of use and simplicity.

## Features

- Clone a PostgreSQL database from one instance to another.
- Simple command-line interface.
- Easy setup with minimal configuration.

## Installation

You can install **pgklone** using the following script:

```bash
curl -sSL https://raw.githubusercontent.com/JuanRulliansyah/pgklone/master/install.sh | bash
```

### Prepare for Windows Users

*Coming Soon*

## Usage

After installation, you can use **pgklone** to clone your PostgreSQL databases interactively. When you run the tool, it will prompt you to enter the URLs for the source and target databases.

### Basic Usage

Run the tool using:

```bash
pgklone
```

You will be prompted to enter the following information:

1. **Source DB URL**: The URL of the database you want to clone from.
2. **Target DB URL**: The URL of the database you want to clone to.

Example interaction:

```plaintext
Enter Source DB URL: postgres://user:password@localhost:5432/source_db
Enter Target DB URL: postgres://user:password@localhost:5432/target_db
```

After providing the URLs, **pgklone** will display the entered URLs and proceed with the cloning process. You will see a success message once the database has been cloned successfully.

```plaintext
Source DB URL: postgres://user:password@localhost:5432/source_db
Target DB URL: postgres://user:password@localhost:5432/target_db
Database cloned successfully!
```

## License

**pgklone** is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- PostgreSQL for providing a powerful and open-source database.
- All contributors for their valuable input and improvements.