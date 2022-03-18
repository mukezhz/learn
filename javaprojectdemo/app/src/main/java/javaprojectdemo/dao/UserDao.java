package javaprojectdemo.dao;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.ArrayList;
import java.util.List;
 
import javaprojectdemo.model.User;
import javaprojectdemo.util.DbUtil;
 
public class UserDao {
 
    private Connection connection;
 
    public UserDao() {
        connection = DbUtil.getConnection();
    }
     
    public UserDao(Connection conn) {
        connection = conn;
    }

    public void createUserTable() {
        try {
            Statement stmt = connection.createStatement();
            String sql = "CREATE TABLE users( " +
                    " username varchar(30) NOT NULL, " +
                    " password varchar(40) NOT NULL, " +
                    " full_name varchar(45) NOT NULL, " +
                    " email varchar(100) DEFAULT NULL, " +
                    " PRIMARY KEY (username));";

            stmt.executeUpdate(sql);
            System.out.println("Created table in given database...");
        } catch (SQLException e) {
            System.out.println("Table already created");
        }
    }

    public void addUser(User user) {
        try {
            PreparedStatement preparedStatement = connection
                    .prepareStatement("insert into users(username,password,full_name,email) values (?, ?, ?, ? )");
            // Parameters start with 1
            preparedStatement.setString(1, user.getUsername());
            preparedStatement.setString(2, user.getPassword());
            preparedStatement.setString(3, user.getFullName());
            preparedStatement.setString(4, user.getEmail());
            preparedStatement.executeUpdate();
            preparedStatement.close();
            System.out.println("User has been added successfully");
        } catch (SQLException e) {
            System.out.println("username already exits");
            // e.printStackTrace();
        }
    }
 
    public void deleteUser(String username) {
        try {
            PreparedStatement preparedStatement = connection
                    .prepareStatement("delete from users where username=?");
            // Parameters start with 1
            preparedStatement.setString(1, username);
            preparedStatement.executeUpdate();
            preparedStatement.close();
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }
 
    public void updateUser(User user) {
        try {
            PreparedStatement preparedStatement = connection
                    .prepareStatement("update users set password=?, full_name=?, email=?" +
                            "where username=?");
            // Parameters start with 1
            preparedStatement.setString(1, user.getPassword());
            preparedStatement.setString(2, user.getFullName());
            preparedStatement.setString(3, user.getEmail());
            preparedStatement.setString(4, user.getUsername());
            preparedStatement.executeUpdate();
            preparedStatement.close();
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }
 
    public List<User> getAllUsers() {
        List<User> users = new ArrayList<User>();
        try {
            Statement statement = connection.createStatement();
            ResultSet rs = statement.executeQuery("select * from users");
            while (rs.next()) {
                User user = new User();
                user.setUsername(rs.getString("username"));
                user.setPassword(rs.getString("password"));
                user.setFullName(rs.getString("full_name"));
                user.setEmail(rs.getString("email"));
                users.add(user);
            }
        } catch (SQLException e) {
            e.printStackTrace();
        }
 
        return users;
    }
 
    public User getUserByUsername(String username) {
        User user = new User();
        try {
            PreparedStatement preparedStatement = connection.
                    prepareStatement("select * from users where username=?");
            preparedStatement.setString(1, username);
            ResultSet rs = preparedStatement.executeQuery();
 
            if (rs.next()) {
                user.setUsername(rs.getString("username"));
                user.setPassword(rs.getString("password"));
                user.setFullName(rs.getString("full_name"));
                user.setEmail(rs.getString("email"));
            }
        } catch (SQLException e) {
            e.printStackTrace();
        }
 
        return user;
    }
}