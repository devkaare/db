<h1>Not for production use and not recommended for large datasets!</h1>
<div>
  <p>This <code>db</code> package provides basic functions for managing users stored in a JSON file.</p>
</div>
<div>
  <p><strong>Important:</strong><br>
    Call <code>LoadCache()</code> at the start of your main function to initialize the user cache, and defer <code>SaveCache()</code> to save changes when the program exits.</p>
</div>
<div>
  <h3>User Struct:</h3>
  <pre><code>
type User struct {
    Id       int    `json:"Id"`
    Username string `json:"Username"`
    Email    string `json:"Email"`
    Password string `json:"Password"`
}
  </code></pre>
  <p>This can be modified by simply downloading the <code>db.go</code> file, placing it in a folder named <code>db</code> inside your prokect and importing it using <code>import "your_project_path/db"</code>.</p>
</div>
<div>
  <h3>Functions:</h3>
  <p><strong>1. DoesFileExist(path string) (bool, error):</strong><br>
    Checks if a file exists at the given path.</p>
  <p><strong>2. LoadCache():</strong><br>
    Loads the user cache from the JSON file. Call this at the start of your main function.</p>
  <p><strong>3. SaveCache():</strong><br>
    Saves the user cache to the JSON file. Defer this at the start of your main function.</p>
  <p><strong>4. AddUser(user User):</strong><br>
    Adds a user to the cache.</p>
  <p><strong>5. GetUserById(id int) User:</strong><br>
    Retrieves a user from the cache by their Id.</p>
  <p><strong>6. DeleteUserById(id int):</strong><br>
    Deletes a user from the cache by their Id.</p>
</div>
