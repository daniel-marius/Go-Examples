import { useEffect, useState } from "react";
import { BrowserRouter, Route } from "react-router-dom";

import "./App.css";
import Nav from "./components/Nav";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";
import endpoint from "./api/endpoint";

interface UserModel {
  name: string;
  email: string;
  password: string;
}

const App = () => {
  const [name, setName] = useState("");

  useEffect(() => {
    (async () => {
      // const response = await fetch("http://localhost:4000/api/user", {
      //   headers: { "Content-Type": "application/json" },
      //   credentials: "include"
      // });
      //
      // const content = await response.json();

      const { data } = await endpoint.get<UserModel>("/user", { withCredentials: true });

      setName(data.name);
    })();
  });

  return (
    <div className="App">
      <BrowserRouter>
        <Nav name={name} setName={setName} />

        <main className="form-signin">
          <Route path="/" exact component={() => <Home name={name} />} />
          <Route path="/login" component={() => <Login setName={setName} />} />
          <Route path="/register" component={Register} />
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
