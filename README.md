# File Transfer App

## Overview

File Transfer App is a simple web application built with Golang that allows users to upload files and then download them from the server. The app is modularized for better maintainability, with separate handlers for different functionalities.

## Features

- **File Upload**: Upload files through a web form.
- **File Download**: Download uploaded files from the server.
- **Modular Structure**: The code is organized into different packages for easy management.

## Project Structure
file-transfer-app/
│
├── main.go
├── handlers/
│   ├── index.go
│   ├── upload.go
│   ├── download.go
│   ├── file.go
│
├── templates/
│   ├── index.html
│   ├── download.html
│
└── uploads/
    └── (uploaded files go here)

