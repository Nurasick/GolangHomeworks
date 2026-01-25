import { useState } from "react"
import { Link } from "react-router-dom";

function RegisterPage() {

    const [email,setEmail] = useState("");
    const [password,setPassword] = useState("");
    const handleSubmit = async (e) => {
        
        e.preventDefault();
        const response = await fetch("http://localhost:8080/auth/register",{
            headers:{
            "Content-Type": "application/json",
            },
            method: "POST",
            body: JSON.stringify({
                email: email,
                password: password,
                role_id: 1,
            }),
        })
        console.log(response)
    };
    return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-900 to-slate-800">
        <div className="w-full max-w-md rounded-2xl bg-white/10 backdrop-blur-lg shadow-xl p-8" >
            <h1 className="text-3xl font-semibold text-white">University Portal</h1>
            <p className="text-sm text-gray-300 mt-2">Sign up to get started</p>
            <div className="mt-6">
            <form className="flex flex-col gap-1" onSubmit={handleSubmit}>
                <label htmlFor="email" className="block text-sm text-gray-300 mb-1">Email</label>
                <input type = "email"
                id = "email" 
                name="email"
                value={email}
                onChange={e => setEmail(e.target.value)}
                className = "w-full rounded-lg bg-white/5 text-white border border-white/10 px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500 transition"
                ></input>
                <div className="mt-6">
                <label htmlFor="password" className="block text-sm text-gray-300 mb-1">Password</label>
                <input type = "password" 
                id = "password" 
                name="password"
                value={password}
                onChange={e => setPassword(e.target.value)}
                className="
                w-full rounded-lg bg-white/5 text-white border border-white/10 px-4 py-2
                focus:outline-none focus:ring-2 focus:ring-indigo-500 transition
                "></input>
            </div>
            <button type = "submit" className="
            mt-6 w-full py-2 rounded-lg
            bg-gradient-to-r from-indigo-500 to-indigo-500
            text-white font-medium
            hover:opacity-90
            hover:scale-[1.02]
            cursor-pointer
            transition
            ">Sign Up</button>
            <p className="text-sm text-gray-400 mt-6 text-center">Already have an account?  
                <Link to="/login" className="text-indigo-400 hover:underline cursor-pointer"> Log in</Link>
            </p>
            </form>
            </div>
        </div>
    </div>
    )
}
export{RegisterPage}