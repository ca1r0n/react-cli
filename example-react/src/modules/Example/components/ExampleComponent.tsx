import React, {useState} from "react";

export function ExampleComponent() {
    const [test, setTest] = useState("test")

    return <div>
        <input value={test} type="text" onChange={(v) => setTest(v.target.value)}/>
        <div>
            {test}
        </div>
    </div>
}