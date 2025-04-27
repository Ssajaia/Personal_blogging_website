package main

const homeTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>My Personal Blog</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        .article { margin-bottom: 30px; border-bottom: 1px solid #eee; padding-bottom: 20px; }
        .article h2 a { color: #333; text-decoration: none; }
        .date { color: #666; font-size: 0.9em; }
        .admin-link { float: right; }
    </style>
</head>
<body>
    <h1>My Personal Blog</h1>
    <a href="/admin/" class="admin-link">Admin Login</a>
    
    {{range .}}
    <div class="article">
        <h2><a href="/article/{{.ID}}">{{.Title}}</a></h2>
        <div class="date">Published on {{.Date}}</div>
        <p>{{printf "%.200s" .Content}}...</p>
        <a href="/article/{{.ID}}">Read more</a>
    </div>
    {{end}}
</body>
</html>
`

const articleTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}} - My Personal Blog</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        .date { color: #666; font-size: 0.9em; }
        .content { line-height: 1.6; }
        .back-link { display: block; margin-top: 20px; }
    </style>
</head>
<body>
    <h1>{{.Title}}</h1>
    <div class="date">Published on {{.Date}}</div>
    <div class="content">{{.Content}}</div>
    <a href="/" class="back-link">← Back to all articles</a>
</body>
</html>
`

const dashboardTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Admin Dashboard</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        .article { margin-bottom: 20px; padding: 15px; border: 1px solid #ddd; border-radius: 5px; }
        .actions a { margin-right: 10px; }
        .add-btn { display: inline-block; margin-bottom: 20px; padding: 8px 15px; background: #4CAF50; color: white; text-decoration: none; border-radius: 4px; }
        .logout { float: right; }
    </style>
</head>
<body>
    <h1>Admin Dashboard <a href="/admin/logout" class="logout">Logout</a></h1>
    <a href="/admin/add" class="add-btn">Add New Article</a>
    
    {{range .}}
    <div class="article">
        <h3>{{.Title}}</h3>
        <div class="date">Published on {{.Date}}</div>
        <div class="actions">
            <a href="/admin/edit/{{.ID}}">Edit</a>
            <a href="/admin/delete/{{.ID}}" onclick="return confirm('Are you sure?')">Delete</a>
            <a href="/article/{{.ID}}">View</a>
        </div>
    </div>
    {{end}}
</body>
</html>
`

const addTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Add New Article</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        form { margin-top: 20px; }
        label { display: block; margin-top: 10px; }
        input[type="text"], textarea { width: 100%; padding: 8px; margin-top: 5px; }
        textarea { height: 300px; }
        button { margin-top: 15px; padding: 8px 15px; background: #4CAF50; color: white; border: none; border-radius: 4px; }
    </style>
</head>
<body>
    <h1>Add New Article</h1>
    <form method="POST">
        <label for="title">Title:</label>
        <input type="text" id="title" name="title" required>
        
        <label for="content">Content:</label>
        <textarea id="content" name="content" required></textarea>
        
        <button type="submit">Save Article</button>
    </form>
</body>
</html>
`

const editTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Edit Article</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        form { margin-top: 20px; }
        label { display: block; margin-top: 10px; }
        input[type="text"], textarea { width: 100%; padding: 8px; margin-top: 5px; }
        textarea { height: 300px; }
        button { margin-top: 15px; padding: 8px 15px; background: #4CAF50; color: white; border: none; border-radius: 4px; }
    </style>
</head>
<body>
    <h1>Edit Article</h1>
    <form method="POST">
        <label for="title">Title:</label>
        <input type="text" id="title" name="title" value="{{.Title}}" required>
        
        <label for="content">Content:</label>
        <textarea id="content" name="content" required>{{.Content}}</textarea>
        
        <button type="submit">Update Article</button>
    </form>
</body>
</html>
`

const loginTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Admin Login</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 400px; margin: 0 auto; padding: 20px; }
        form { margin-top: 20px; }
        label { display: block; margin-top: 10px; }
        input[type="text"], input[type="password"] { width: 100%; padding: 8px; margin-top: 5px; }
        button { margin-top: 15px; padding: 8px 15px; background: #4CAF50; color: white; border: none; border-radius: 4px; }
    </style>
</head>
<body>
    <h1>Admin Login</h1>
    <form method="POST">
        <label for="username">Username:</label>
        <input type="text" id="username" name="username" required>
        
        <label for="password">Password:</label>
        <input type="password" id="password" name="password" required>
        
        <button type="submit">Login</button>
    </form>
</body>
</html>
`