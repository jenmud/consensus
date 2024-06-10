CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    created_at DATETIME NOT NULL DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%S', 'NOW', 'UTC')),
    updated_at DATETIME NOT NULL DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%S', 'NOW', 'UTC')),
    email TEXT NOT NULL UNIQUE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password TEXT NOT NULL, -- NOTE: we do not store the password in plain text, so hash it using the bcrypt algorithm.
    role TEXT CHECK (role IN ('admin', 'user')) NOT NULL DEFAULT 'user'
);

CREATE TRIGGER IF NOT EXISTS update_user
AFTER UPDATE ON users
BEGIN              
    UPDATE users SET updated_at = STRFTIME('%Y-%m-%d %H:%M:%S', 'NOW', 'UTC')
    WHERE users.id = NEW.id;                   
end;                                     
                                         
CREATE TABLE IF NOT EXISTS epic (        
    id INTEGER PRIMARY KEY,              
    created_at DATETIME NOT NULL DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%S', 'NOW', 'UTC')),
    updated_at DATETIME NOT NULL DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%S', 'NOW', 'UTC')),
    name TEXT NOT NULL,                  
    description TEXT,                    
    user_id INTEGER NOT NULL,            
    FOREIGN KEY (user_id) REFERENCES users(id)
);                                       

CREATE TRIGGER IF NOT EXISTS update_epic
AFTER UPDATE ON epic
BEGIN              
    UPDATE epic SET updated_at = STRFTIME('%Y-%m-%d %H:%M:%S', 'NOW', 'UTC')
    WHERE epic.id = NEW.id;                   
end;                                     
                                         
CREATE TABLE IF NOT EXISTS project (     
    id INTEGER PRIMARY KEY,              
    created_at DATETIME NOT NULL DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%S', 'NOW', 'UTC')),
    updated_at DATETIME NOT NULL DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%S', 'NOW', 'UTC')),
    name TEXT NOT NULL,                  
    description TEXT,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);                           

CREATE TRIGGER IF NOT EXISTS update_project
AFTER UPDATE ON project
BEGIN              
    UPDATE project SET updated_at = STRFTIME('%Y-%m-%d %H:%M:%S', 'NOW', 'UTC')
    WHERE project.id = NEW.id;                   
end;                                     