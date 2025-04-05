import "./App.css"
import { RouterProvider } from "react-router-dom";
import { browserRouter } from "./router";

export function App() {
  return (
    <RouterProvider router={browserRouter} />
  )
}
