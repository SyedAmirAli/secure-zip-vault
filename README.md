# 🏗 SecureZipVault - Project Architecture

## 📌 Project Overview
**SecureZipVault** is a high-security backend microservice built in **Go (Golang)** that:
- **Automatically backs up project files & MySQL database** as a ZIP file.
- **Stores backups on Google Drive** daily.
- **Serves project ZIP downloads** via a password-protected web interface.
- **Ensures high security, efficiency, and performance.**

## 🚀 Technology Stack
- **Backend**: Go (Golang) ✅
- **Database**: MySQL ✅
- **Frontend**: React.js / Vue.js ✅
- **Storage**: Google Drive API ✅
- **Authentication**: JWT (JSON Web Tokens) ✅
- **Encryption**: AES-256 ✅
- **Deployment**: VPS (Ubuntu/Debian) ✅
- **Task Scheduling**: Cron Jobs ✅

## 📂 Project Structure
```
SecureZipVault/
│-- backend/
│   ├── main.go              # Main server entry point
│   ├── config/
│   │   ├── config.go        # App configuration
│   ├── handlers/
│   │   ├── auth.go          # Authentication handling
│   │   ├── download.go      # ZIP file download handler
│   ├── services/
│   │   ├── backup.go        # MySQL & File backup logic
│   │   ├── gdrive.go        # Google Drive API integration
│   ├── utils/
│   │   ├── encrypt.go       # AES-256 encryption for security
│   │   ├── zip.go           # File compression logic
│   ├── routes.go            # API routes
│-- frontend/
│   ├── src/
│   │   ├── App.js           # React/Vue entry point
│   │   ├── components/
│   │   │   ├── LoginForm.js # Password input UI
│   │   │   ├── Download.js  # File download UI
│   ├── package.json         # Frontend dependencies
│-- .env                     # Environment variables
│-- README.md                # Project documentation
```

## 🔐 Security Considerations
1. **Authentication**: JWT-based user login with **bcrypt password hashing**.
2. **Encryption**:
   - MySQL dumps & ZIP files are **AES-256 encrypted** before storage.
   - Secure password storage using **bcrypt**.
3. **Access Control**:
   - IP-based rate-limiting to prevent brute force attacks.
   - Download requests require an **authorization token**.
4. **Secure File Handling**:
   - Use **Go’s os & io packages** for efficient file management.
   - Ensure **temporary files auto-delete** after a defined retention period.
5. **HTTPS Support**: Enforce **TLS (SSL) encryption** for secure API communication.

## ⚡ API Endpoints
### **1️⃣ User Authentication**
```http
POST /api/auth/login
```
- **Body**: `{ "password": "your_secure_password" }`
- **Response**: `{ "token": "JWT_ACCESS_TOKEN" }`

### **2️⃣ Download Project Backup**
```http
GET /api/download
Authorization: Bearer <JWT_TOKEN>
```
- **Response**: Initiates a secure ZIP file download.

### **3️⃣ Google Drive Backup Upload**
```http
POST /api/backup/upload
Authorization: Bearer <JWT_TOKEN>
```
- **Response**: `{ "status": "Backup uploaded successfully" }`

## ⏳ Task Scheduling
- **Daily Auto Backup** using **cron jobs** (`backup.go`):
  - Exports MySQL database.
  - Zips database + project files.
  - Encrypts ZIP before upload.
  - Uploads to Google Drive.

## 🚀 Deployment Guide
### **1️⃣ VPS Setup (Ubuntu)**
```bash
sudo apt update && sudo apt install golang -y
```

### **2️⃣ Clone the Project & Build**
```bash
git clone https://github.com/your-repo/SecureZipVault.git
cd SecureZipVault/backend
go build -o securezipvault
```

### **3️⃣ Run the Server**
```bash
./securezipvault
```

## 📌 Future Enhancements
- [ ] **WebSocket support** for real-time download status.
- [ ] **Multi-user authentication** with role-based access.
- [ ] **AWS S3 integration** as an alternative backup option.

---
🎯 **SecureZipVault is designed to be a high-performance, secure, and automated backup system for your projects.** 🚀

