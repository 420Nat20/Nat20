import React from "react";
import Head from "next/head";
import "../styles/tailwind.css";
import { Auth0Provider } from "@auth0/auth0-react";
import { Provider } from "next-auth/client";

function MyApp({ Component, pageProps }) {
  return (
    <Provider session={pageProps.session}>
      <Head>
        <title>Nat20</title>
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width"
        />
      </Head>
      <Component {...pageProps} />
    </Provider>
  );
}

export default MyApp;
