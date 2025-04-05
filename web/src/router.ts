import { createBrowserRouter, RouteObject } from "react-router-dom";
import Layout from "./components/Layout";
import Home from "./pages/Home";

export const routes: RouteObject[] = [
    {
        Component: Layout,
        children: [
            // Home page
            { index: true, Component: Home }
        ]
    }
]

export const browserRouter = createBrowserRouter(routes)