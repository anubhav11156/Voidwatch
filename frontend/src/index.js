import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import ErrorPage from './components/ErrorPage';
import Home from './components/Home';
import reportWebVitals from './reportWebVitals';
import {
  createBrowserRouter,
  RouterProvider,
  Route,
  Link,
} from 'react-router-dom'
import Movies from './components/Movies';
import Genres from './components/Genres';
import Movie from './components/Movie';
import ManageCatlouge from './components/ManageCatlouge';
import GraphQl from './components/GraphQl';
import Login from './components/Login';
import EditMovie from './components/EditMovie';

const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {index: true, element: <Home />},
      {
        path:'/movies',
        element: <Movies />
      },
      {
        path:'/movies/:id',
        element: <Movie />
      },
      {
        path:'/genres',
        element: <Genres />
      },
      {
        path:'/admin/movie/0',
        element: <EditMovie />
      },
      {
        path:'/manage-catalouge',
        element: <ManageCatlouge />
      },
      {
        path:'/graphql',
        element: <GraphQl />
      },
      {
        path:'/login',
        element: <Login />
      },
    ]
  }
]) 


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
