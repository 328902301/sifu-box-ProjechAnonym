import { Route, Routes } from "react-router-dom";

import LoginPage from "./pages/login";
import HomePage from "./pages/home";
import SettingPages from "./pages/setting";
function App() {
  return (
    <div>
      <Routes>
        <Route element={<LoginPage />} path="/" />
        <Route element={<HomePage />} path="/home" />
        <Route element={<SettingPages />} path="/setting" />
      </Routes>
    </div>
  );
}

export default App;
