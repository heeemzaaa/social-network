# Social Network Application

A full-stack **Social Network Application** built with a **Go backend** and a **Next.js frontend**. The project demonstrates modern web development practices including clean architecture, database migrations, real-time communication, authentication, and containerization with Docker.

---

##  Features

### 🔹 Authentication & Profiles
- User registration & login with sessions and cookies
- Profile pages (public or private)
- Ability to update profile info (avatar, nickname, about me, etc.)
- Follow/unfollow functionality with support for **follow requests** on private accounts

### 🔹 Posts & Groups
- Create posts (with images/GIFs) and comments
- Control post privacy: **public**, **followers only**, or **custom private**
- Group creation with invitations and join requests
- Group posts, comments, and group chat
- Group events with RSVP options (Going / Not Going)

### 🔹 Real-Time Features
- **Private chat** between users (via WebSockets)
- **Group chat** for all members
- **Notifications system** for:
  - Follow requests
  - Group invitations
  - Group join requests
  - Group event creation

### 🔹 Frontend (Next.js)
- Client-side rendering and routing
- **Group routing** for organized navigation
- **Private routes** to protect authenticated pages
- Clean and responsive UI

### 🔹 Backend (Go)
- **Three-layer architecture**:
  - **Repository** → database access
  - **Service** → business logic
  - **Handler** → HTTP handlers
- **Database migrations** with `golang-migrate`
- **SQLite** database integration
- **WebSockets** for chat and notifications
- **Goroutines & Channels** for concurrency (chat & notifications)
- Data race handling to ensure safe concurrent access

### 🔹 DevOps
- **Dockerized backend & frontend**
- Separate containers for frontend and backend
- Ready for deployment with exposed ports and environment configs

---

##  Project Structure

### Frontend (Next.js)
```
frontend
├── app
│   ├── _actions
│   │   ├── group.js
│   │   ├── groupPosts.js
│   │   ├── posts.js
│   │   └── user.js
│   ├── (auth)
│   │   ├── login
│   │   │   ├── loginForm.jsx
│   │   │   └── page.jsx
│   │   └── register
│   │       ├── page.jsx
│   │       └── registerForm.jsx
│   ├── _components
│   │   ├── button.jsx
│   │   ├── components.css
│   │   ├── logo.jsx
│   │   └── subimtButton.jsx
│   ├── global.css
│   ├── _hooks
│   │   ├── useDebouncedCallback.js
│   │   ├── useDebounce.js
│   │   └── useInfiniteScroll.js
│   ├── layout.jsx
│   ├── (main)
│   │   ├── chat
│   │   │   ├── chat.css
│   │   │   ├── _components
│   │   │   │   ├── chat_components.css
│   │   │   │   ├── fetchMessages.jsx
│   │   │   │   ├── group_list.jsx
│   │   │   │   └── userList.jsx
│   │   │   └── page.jsx
│   │   ├── _components
│   │   │   ├── avatar.jsx
│   │   │   ├── comments
│   │   │   │   ├── commentsCard.jsx
│   │   │   │   ├── commentsContainer.jsx
│   │   │   │   ├── comments.css
│   │   │   │   ├── commentsFooter.jsx
│   │   │   │   └── comments.jsx
│   │   │   ├── components.css
│   │   │   ├── error.jsx
│   │   │   ├── header.jsx
│   │   │   ├── loader.jsx
│   │   │   ├── modal
│   │   │   │   ├── modal.css
│   │   │   │   └── modal.jsx
│   │   │   ├── navigation.jsx
│   │   │   ├── notifications
│   │   │   │   ├── NotificationsContainer.jsx
│   │   │   │   └── styles.css
│   │   │   ├── posts
│   │   │   │   ├── createPost.jsx
│   │   │   │   ├── postCard.jsx
│   │   │   │   ├── postCardList.jsx
│   │   │   │   ├── postsContainer.jsx
│   │   │   │   └── style.css
│   │   │   ├── tab
│   │   │   │   ├── style.css
│   │   │   │   ├── tabContent.jsx
│   │   │   │   ├── tab.jsx
│   │   │   │   └── tabs.jsx
│   │   │   └── tag.jsx
│   │   ├── _context
│   │   │   ├── ModalContext.js
│   │   │   ├── NotificationContext.jsx
│   │   │   └── userContext.js
│   │   ├── groups
│   │   │   ├── _components
│   │   │   │   ├── createEventForm.jsx
│   │   │   │   ├── createGroupForm.jsx
│   │   │   │   ├── createPostForm.jsx
│   │   │   │   ├── groupCard.jsx
│   │   │   │   ├── groupCardList.jsx
│   │   │   │   ├── groupEventCard.jsx
│   │   │   │   ├── groupEventCardList.jsx
│   │   │   │   ├── groupPostCardList.jsx
│   │   │   │   ├── inviteFriendsForm.jsx
│   │   │   │   ├── style.css
│   │   │   │   └── userCard.jsx
│   │   │   ├── [id]
│   │   │   │   ├── page.jsx
│   │   │   │   └── style.css
│   │   │   └── page.jsx
│   │   ├── layout.jsx
│   │   ├── _lib
│   │   │   └── webSocket.jsx
│   │   ├── page.jsx
│   │   └── profile
│   │       └── [id]
│   │           ├── _components
│   │           │   ├── profileConnections
│   │           │   │   ├── followerCard.jsx
│   │           │   │   ├── followersContainer.jsx
│   │           │   │   └── followers.css
│   │           │   ├── profileData
│   │           │   │   ├── abouUser.jsx
│   │           │   │   ├── profile.css
│   │           │   │   └── userInfo.jsx
│   │           │   └── profilePosts
│   │           │       ├── posts-profile.css
│   │           │       └── userPosts.jsx
│   │           └── page.jsx
│   ├── page.module.css
│   └── _utils
│   └── time.js
|____
```

### Backend (Go)
```
backend
├── database
│   ├── forum.db
│   ├── migrations
│   │   ├── 000001_create_users_table.down.sql
│   │   ├── 000001_create_users_table.up.sql
│   │   ├── ... (other migration files)
│   │   └── 000018_create_notifications_table.up.sql
│   └── sqlite
│       └── sqlite.go
├── go.Dockerfile
├── go.mod
├── go.sum
├── handlers
│   ├── auth
│   ├── chat
│   ├── group
│   ├── notification
│   ├── post
│   ├── profile
│   └── user
├── middleware
├── models
├── repositories
|   ├── auth
│   ├── chat
│   ├── group
│   ├── notification
│   ├── post
│   ├── profile
│   └── user
├── routes
├── services
|   ├── auth
│   ├── chat
│   ├── group
│   ├── notification
│   ├── post
│   ├── profile
│   └── user
├── storage
|── main.go
└── utils
```

---

## 🛠️ Tech Stack

**Frontend:**  
- Next.js (React Framework)  
- HTML, CSS, JavaScript

**Backend:**  
- Go (Golang)
- Gorilla WebSocket
- SQLite
- Golang-Migrate
- bcrypt (password hashing)
- UUID

**Infrastructure:**  
- Docker (multi-container setup)

---

##  Installation & Setup

### 1 Clone the Repository
```bash
git clone https://github.com/heeemzaaa/social-network.git
cd social-network
```

### 2 Run with Docker
Ensure Docker is installed, then run:
```bash
docker-compose up --build
```

This will spin up two containers:
- **Backend (Go server)**
- **Frontend (Next.js app)**

### 3 Access the App
- Frontend: [http://localhost:3000](http://localhost:3000)
- Backend API: [http://localhost:8080](http://localhost:8080)

---

##  Learning Outcomes
This project was designed to cover a broad range of backend and frontend concepts:

- Authentication (sessions & cookies)
- SQL database design & migrations
- Concurrency in Go (goroutines & channels)
- Data race handling
- Real-time communication (WebSockets)
- Docker containerization
- Routing (frontend & backend)
- Clean architecture principles

---

##  Authors
- [Hamza ELKHAWLANI](https://github.com/heeemzaaa)
- [uma-oo](https://github.com/uma-oo)
- [Ayoub NACHTI](https://github.com/DarkMethoss)
- [Ibrahim FARES](https://github.com/Faresbrahiim)
- [Ahmed Amine MELLAGUI](https://github.com/Mellagui)
- [Yassine FAWZI](https://github.com/yassinefawzi)

