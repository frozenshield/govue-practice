import { PublicClientApplication } from "@azure/msal-browser";

export const msalConfig = {
    auth: {
        clientId: import.meta.env.VITE_AZURE_CLIENT_ID,
        authority: import.meta.env.VITE_AZURE_AUTHORITY,
        redirectUri: 'http://localhost:5173/',
        postLogoutRedirectUri: 'http://localhost:5173/',
    },
    cache: {
        cacheLocation: "localStorage", // This configures where your cache will be stored
        storeAuthStateInCookie: true, // Set this to "true" if you are having issues on IE11 or Edge
    }
}

export const loginRequest = {
    scopes: ['openid', 'profile', 'User.Read']
}

export const msalInstance = new PublicClientApplication(msalConfig);

