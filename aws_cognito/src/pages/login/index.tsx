"use client";

import axios from "axios";
import styles from "../../app/page.module.css";
import { Inter } from "next/font/google";
import { redirect, useSearchParams } from "next/navigation";
import { useRouter } from 'next/navigation'
import { useEffect, useState } from "react";

const inter = Inter({ subsets: ["latin"] });

export default function Login() {
  const searchParams = useSearchParams();
  const code = searchParams?.get("code");
  const [token, setToken] = useState("");
  const router = useRouter()

  useEffect(() => {
    if (!code) return;
    console.log(code);
    const clientID = process.env.NEXT_PUBLIC_COGINTO_CLIENT_ID || "";
    const clientSecret = process.env.NEXT_PUBLIC_COGNITO_CLIENT_SECRET || "";
    const cognitoDomain = process.env.NEXT_PUBLIC_COGNITO_DOMAIN || "";
    const credentials = `${clientID}:${clientSecret}`;
    const base64Credentials = Buffer.from(credentials).toString("base64");
    const basicAuthorization = `Basic ${base64Credentials}`;
    const headers = {
      "Content-Type": "application/x-www-form-urlencoded",
      Authorization: basicAuthorization,
    };
    const data = new URLSearchParams();
    let token = localStorage.getItem("access_token") || "";
    if (token) {
      setToken(token);
      return;
    }
    data.append("grant_type", "authorization_code");
    data.append("client_id", clientID);
    data.append("code", code);
    data.append("redirect_uri", "http://localhost:3000/login");
    axios
      .post(`${cognitoDomain}/oauth2/token`, data, { headers })
      .then((res) => {
        if (res.status != 200) return;
        token = res?.data?.access_token;
        console.log(res);
        localStorage.setItem("access_token", token);
        localStorage.setItem("id_token", res?.data?.id_token);
        setToken(token);
      });
  }, [code]);

  // if (token) router.push("/");

  return  <div className={styles.center}>
  <a
    href={`https://test-yomiuri-dev.auth.ap-northeast-1.amazoncognito.com/oauth2/authorize?client_id=3m9do130qvu8sbe2104n0u7qe&response_type=code&scope=email+openid+phone&redirect_uri=http%3A%2F%2Flocalhost%3A3000%2Flogin`}
    className={styles.card}
    // target="_blank"
    rel="noopener noreferrer"
  >
    <h2 className={inter.className}>
      Login <span>-&gt;</span>
    </h2>
    <p className={inter.className}>Cognito</p>
  </a>
</div>;
}
