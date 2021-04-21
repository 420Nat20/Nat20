import NextAuth from "next-auth";
import Providers from "next-auth/providers";

export default NextAuth({
  // Configure one or more authentication providers
  providers: [
    {
      id: "discord",
      name: "Discord",
      type: "oauth",
      version: "2.0",
      scope: "identify email guilds guilds.join",
      params: { grant_type: "authorization_code" },
      accessTokenUrl: "https://discord.com/api/oauth2/token",
      authorizationUrl:
        "https://discord.com/api/oauth2/authorize?response_type=code&prompt=none",
      profileUrl: "https://discord.com/api/users/@me",
      profile(profile) {
        if (profile.avatar === null) {
          const defaultAvatarNumber = parseInt(profile.discriminator) % 5;
          profile.image_url = `https://cdn.discordapp.com/embed/avatars/${defaultAvatarNumber}.png`;
        } else {
          const format = profile.avatar.startsWith("a_") ? "gif" : "png";
          profile.image_url = `https://cdn.discordapp.com/avatars/${profile.id}/${profile.avatar}.${format}`;
        }
        return {
          id: profile.id,
          name: profile.username,
          image: profile.image_url,
          email: profile.email,
        };
      },
      clientId: process.env.DISCORD_CLIENT_ID,
      clientSecret: process.env.DISCORD_SECRET,
      callbacks: {
        async signIn(user, account, profile) {
          // call the Nat20 API and check if it's a new user
          return true;
        },
      },
    },
  ],
  pages: {
    signIn: "/auth/login",
  },
});
