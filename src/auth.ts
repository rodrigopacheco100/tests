import { betterAuth } from "better-auth";
import { drizzleAdapter } from "better-auth/adapters/drizzle";
import { bearer, jwt, openAPI } from "better-auth/plugins";
import { db } from "./database";

export const auth = betterAuth({
	database: drizzleAdapter(db, {
		provider: "sqlite", // or "pg" or "mysql"
	}),
	emailAndPassword: {
		enabled: true,
	},
	plugins: [openAPI(), bearer(), jwt()],
});
