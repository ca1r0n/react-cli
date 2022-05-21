import React from "react";
import {App} from "./App";
import {createRoot} from "react-dom/client";

const app = document.getElementById("app")

if (app) {
    createRoot(app).render(<App />)
}
