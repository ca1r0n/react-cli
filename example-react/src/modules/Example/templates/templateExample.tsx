import React from "react";
import {ExampleInterface} from "../containers/containerExample";
import classNames from "classnames";

export function TemplateExample(c: ExampleInterface) {
    return <div className={classNames("example")}>
        {c.state.example}
    </div>
}