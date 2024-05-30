<h1><strong>WARNING:</strong> This is not meant for production use, and is not recommended when dealing with large amounts of data!</h1>

<div>
  <h2>Package: db</h2>
  <p>Package <code>db</code> provides functions for interacting with a user database stored in a JSON file.</p>
</div>
<div>
  <p><strong>Important Note:</strong><br>
   - <code>FillCache()</code> and <code>defer SaveCache()</code> functions **should be placed at the top of your main function.** <code>FillCache()</code> initializes the cache by reading from the JSON file, and <code>SaveCache()</code> ensures that any changes made to the cache are saved back to the JSON file when the program exits.</p>
</div>
<div>
  <h3>Functions:</h3>

  <p><strong>1. DoesFileExist(path string) (bool, error):</strong><br>
   - Checks if a file exists at the specified path.<br>
   - Returns <code>true</code> and <code>nil</code> error if the file exists, <code>false</code> and <code>nil</code> error if it doesn't, and <code>true</code> with an error if any other error occurs.</p>

  <p><strong>2. FillCache():</strong><br>
   - Initializes the user cache by reading from the JSON file.<br>
   - Should be called at the beginning of the main function.</p>

  <p><strong>3. SaveCache():</strong><br>
   - Saves the user cache to the JSON file.<br>
   - Should be deferred at the beginning of the main function.</p>

  <p><strong>4. SaveUser(user User):</strong><br>
   - Adds a user to the cache.</p>

  <p><strong>5. GetUsers() []User:</strong><br>
   - Retrieves all users from the cache.</p>

  <p><strong>6. GetUserById(id int) User:</strong><br>
   - Retrieves a user from the cache by their ID.</p>

  <p><strong>7. DeleteUserById(id int):</strong><br>
   - Removes a user from the cache by their ID.</p>
</div>
