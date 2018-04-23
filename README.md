# Another Go PDF library

*****

Another PDF library, the idea is to easily create PDF files following a TOML (originally JSON) definition document.

There are three important parts used to create a document:
1. **The template**. This file describes the expected fields and the type of those fields.
2. **The layout**. One template can have several layouts, meaning that you can create several different outputs from the same data and for the same purpose.
3. **The data**. This is the only document that a regular user needs to create, the data will be used to replace placeholders in the templates.

The project includes a library of common templates used for common tasks, like creating a CV, a invoice, a receipt or a model release. Also there is a layout library, at this time there are not many layouts but I hope to generate a few myself and I will be happy to include templates and layouts created by other users.

## Usage
Right now the only way to use this tool is by cloning the repo into your go path and build the project by yourself. I am still learning Go, and one of my priorities once a functional version is ready will be to make easier to install and use the tool.
```
pdf2go -in=someFile.toml -out=somePDFFile.pdf -template=templateName
```
