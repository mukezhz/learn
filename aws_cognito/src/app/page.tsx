"use client";

import { Inter } from "next/font/google";
import styles from "./page.module.css";
import { useEffect, useState } from "react";
import axios from "axios";
import { useRouter } from "next/navigation";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const router = useRouter();

  useEffect(() => {
    let token = localStorage.getItem("access_token") || "";
    if (!token) return;
    const cognitoDomain = process.env.NEXT_PUBLIC_COGNITO_DOMAIN || "";
    console.log(token);
    const userInfoHeaders = {
      Authorization: "Bearer " + token,
    };
    axios
      .get(`${cognitoDomain}/oauth2/userInfo`, {
        headers: userInfoHeaders,
      })
      .then((userInfo) => {
        if (userInfo.status != 200) return;
        setName(userInfo.data?.username);
        setEmail(userInfo.data?.email);
      });
  }, []);

  return (
    <main className={styles.main}>
      <div className={styles.description}></div>
      <h2 className={inter.className}>Welcome to Cognito!</h2>
      {name && email ? (
        <>
          <h2 className={inter.className}>{name}</h2>
          <p className={inter.className}>{email}</p>

            <button onClick={() => {localStorage.removeItem("access_token"); router.push("/login")}}>Logout</button>

        </>
      ) : (
        <>
        </>
      )}
    </main>
  );
}
