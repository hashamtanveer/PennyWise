import { Route } from 'react-router-dom';
import Login from '../components/Login';
import Dashboard from '../components/Dashboard';
import Form from '../components/Form';
import Signup from '../components/Signup';
import LoginForm from '../components/LoginForm';


const generalRoutes = (
    <Route key="home" path="/">
   <Route
   path=""
   element={<Login/>
   }
   />
    <Route
   path="/dashboard"
   element={<Dashboard/>
   }
   />
    <Route
   path="/form"
   element={<Form/>
   }
   />
    <Route
   path="/signup"
   element={<Signup/>
   }
   />
    <Route
   path="/login"
   element={<LoginForm/>
   }
   />
   </Route>
   );
export default generalRoutes;