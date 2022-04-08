import "reflect-metadata"
import { DataSource } from "typeorm"

export const AppDataSource = new DataSource({
  type: "postgres",
  host: "db",
  port: 5432,
  username: "rroot",
  password: "root",
  database: "nest",
  synchronize: true,
  logging: true
})