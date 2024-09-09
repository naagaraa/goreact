import React from 'react'
import { Link } from 'react-router-dom';

export default function Navbar() {
    return (
        <>
            <nav class="bg-white shadow dark:bg-gray-800">
                <div class="container flex items-center justify-center p-6 mx-auto text-gray-600 capitalize dark:text-gray-300">

                    <Link className='text-gray-800 transition-colors duration-300 transform dark:text-gray-200 border-b-2 border-blue-500 mx-1.5 sm:mx-6' to="/">Home</Link>

                    <Link className='text-gray-800 transition-colors duration-300 transform dark:text-gray-200 border-b-2 border-blue-500 mx-1.5 sm:mx-6' to="/about">About</Link>


                </div>
            </nav>
        </>
    )
}
