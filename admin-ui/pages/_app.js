import React from "react";
import Head from "next/head";
import "../styles/tailwind.css";
import { Auth0Provider } from "@auth0/auth0-react";
function MyApp({ Component, pageProps }) {
  return (
    <Auth0Provider
      domain="mfarver.us.auth0.com"
      clientId="OtnUwjaPrcHYXsqMP9599k8ZskfQz0Bw"
      redirectUri="http://localhost:3000/"
    >
      <Head>
        <title>Nat20</title>
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width"
        />
      </Head>
      <Component {...pageProps} />
    </Auth0Provider>
  );
}

export default MyApp;
