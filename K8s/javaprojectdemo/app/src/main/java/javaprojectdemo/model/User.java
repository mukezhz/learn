package javaprojectdemo.model;

public class User {
 
    private String username;
    private String password;
    private String fullName;
    private String email;
     
    public String getUsername() {
        return username;
    }
    public void setUsername(String username) {
        this.username = username;
    }
    public String getPassword() {
        return password;
    }
    public void setPassword(String password) {
        this.password = password;
    }
    public String getFullName() {
        return fullName;
    }
    public void setFullName(String fullName) {
        this.fullName = fullName;
    }
    public String getEmail() {
        return email;
    }
    public void setEmail(String email) {
        this.email = email;
    }
    @Override
    public String toString() {
        return "User [username=" + username + ", password=" + password + ", fullName=" + fullName + ", email=" + email
                + "]";
    }
     
}