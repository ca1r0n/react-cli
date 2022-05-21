import React from "react";
import {Link, Route, Routes} from "react-router-dom";
import {Example} from "./modules/Example/containers/containerExample";

export function AppRouter() {
    return <>
        <Link to={"/"}>Test react cli</Link>
        <Routes>
            <Route path={"/"} element={<Example />} />
        </Routes>
    </>
}