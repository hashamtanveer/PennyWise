import {
    createBrowserRouter,
    createRoutesFromElements,
} from 'react-router-dom';
import generalRoutes from './routes/routes';

const router = createBrowserRouter(
    createRoutesFromElements([
        generalRoutes
    ])
);
export default router;
