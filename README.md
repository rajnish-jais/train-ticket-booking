## **Train Ticket Booking System (gRPC + Golang)**

This is a gRPC-based train ticket booking system built with Golang.
The system allows users to book tickets, retrieve receipts, view allocated seats,
remove users, and modify seat assignments dynamically. 
It follows SOLID principles and uses an adapter pattern to separate storage logic.

### 📌 Features
* ✔ Purchase a Train Ticket (Auto-allocated seat in Section A or B)
* ✔ View Receipt Details (From, To, User, Price, Seat, Section)
* ✔ Get Users by Section (Lists all users and seats in a given section)
* ✔ Remove a User (Frees the seat and removes the user from the system)
* ✔ Modify a User's Seat (Change a user's seat if available)

### 🛠 Tech Stack
* Golang (Backend implementation)
* gRPC (Efficient communication protocol)
* Protocol Buffers (Defines structured messages)
* SOLID Principles (Scalability & Maintainability)
* Adapter Pattern (Storage abstraction with an in-memory adapter)

### **⚙️ Installation & Setup**

1️⃣ Clone the Repository
* git clone https://github.com/rajnish-jais/train-ticket-booking.git
* cd train-ticket-booking

2️⃣ Install Dependencies: 
* Ensure you have Golang installed. Then, install required dependencies:
* make install


3️⃣ Generate gRPC Code (if needed)
* If modifying the .proto file, regenerate gRPC code:
* make proto

### **🚀 Running the Application**
**Start the gRPC Server**

* make run-server #Server will start on the configured port (default: 50051).

**Run the gRPC Client**

* make run-client

### 🛠 Testing
**Run unit tests using:**

* make test
* make coverage

**Clean up binaries and generated files**

* make clean


