import { createBrowserRouter } from "react-router";
import { RouterProvider } from "react-router/dom";
import { LoginPage } from './pages/login/page';
import { RegisterPage } from './pages/register/page';
import { Navigate} from "react-router-dom"

const router = createBrowserRouter([
  {
    path: "/",
    element: <Navigate to="/login"/>
  },

  {
    path: "/login",
    element: <LoginPage />,
  },

  {
    path: "/register",
    element: <RegisterPage />,
  }
]);

function App() {
  return <RouterProvider router={router}/>;
}

export default App
