# ⚡ Electricity Billing System (Golang )

## 🚀 Features
- Tariff Management (NTPL, Secunderabad)
- Monthly Billing Calculation
- Progressive Slab Billing
- First & Last Reading Logic
- Consumer-wise Billing
- React Frontend UI

---

## 🧠 Backend (Golang)
- Gin Framework
- GORM ORM
- MySQL Database

---

## 💻 Frontend (React)
- Dropdown selection
- Date-based billing
- Bill breakdown display

---

## 📊 Billing Logic
- Uses FIRST reading of month
- Uses LAST reading of month
- Units = Last - First
- Slab-wise progressive calculation

---

## ⚙️ Setup

### Backend
```bash
cd backend
go mod tidy
go run main.go
