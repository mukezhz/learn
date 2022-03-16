package javaprojectdemo.util;

import java.io.IOException;
import java.io.InputStream;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.util.Properties;
 
public class DbUtil {
     
    private static Connection connection = null;
    private static Properties prop = null;
 
    public static String getProperty(String key) {
        if (prop != null) {
            return prop.getProperty(key);
             
        } else {
            Properties prop = new Properties();
            InputStream inputStream = DbUtil.class.getClassLoader().getResourceAsStream("database.properties");
            try {
                prop.load(inputStream);
                return prop.getProperty(key);
            } catch (IOException e) {
                e.printStackTrace();
                return null;
            }
        }
    }
     
    public static Connection getConnection() {
        if (connection != null)
            return connection;
        else {
            try {
                String driver = getProperty("driver");
                String url = getProperty("url");
                String user = getProperty("user");
                String password = getProperty("password");
                Class.forName(driver);
                connection = DriverManager.getConnection(url, user, password);
//                connection.setAutoCommit(false);
            } catch (ClassNotFoundException e) {
                e.printStackTrace();
            } catch (SQLException e) {
                e.printStackTrace();
            } 
            return connection;
        }
 
    }
 
}