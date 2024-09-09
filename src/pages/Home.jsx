import { Outlet } from 'react-router-dom';
import Navbar from '../components/Navbar';

function Home() {
  return (
    <>
      <Navbar />
      <h1 className="text-3xl font-bold text-center mt-10">
        home pages
      </h1>
      <Outlet />
    </>
  )
}

export default Home
