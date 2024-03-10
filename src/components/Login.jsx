import { Button } from 'antd';
import React from 'react';
import { useNavigate } from 'react-router-dom';

const Login = () => {
	let navigate = useNavigate();

	const navigateToLogin = () => {
	  navigate('/login');
	};
	const navigateToSignup = () => {
		navigate('/signup');
	  };
	return (
		<div className="h-screen">
			{/* Header */}
			<div className="w-full h-16 border-2 border-b-gray-300 flex justify-between items-center px-5">
				<p className="font-semibold text-xl">PennyWise</p>
				<Button type="primary" className="bg-sky-500 font-semibold" onClick={navigateToSignup} >
					Sign Up
				</Button>
			</div>

			{/* Body */}
			<div className="h-[90vh] w-full flex justify-center items-center">
				<div className="w-[400px] h-[700px] flex flex-col items-center justify-center">
					<div className="h-[400px] w-[400px] flex justify-center items-end">
						<img
							src="https://img.freepik.com/free-vector/geometric-gradient-futuristic-background_23-2149116406.jpg?w=1060&t=st=1710091335~exp=1710091935~hmac=a5aec9a8d68aaa6902737115a5897ab7a2da3b384646372dc3f054e6fe9369f9"
							className="h-[400px] w-[400px] rounded-2xl z-0 absolute"
						/>
						<div className="z-10 p-10">
							<p className="text-white text-5xl font-semibold">
								Welcome To
							</p>
							<p className="text-white text-5xl font-semibold mt-3">
								PennyWise
							</p>
							<p className="text-white text-sm font-semibold mt-3">
								Track your income and expenses and manage your
								budget
							</p>
							<Button
								type="primary"
								className="bg-sky-500 font-semibold mt-12"
								onClick={navigateToLogin}
							>
								Login
							</Button>
						</div>
					</div>
					<p className="text-gray-500 mt-6">or</p>
					<Button
						type="primary"
						className="bg-gray-300 text-black font-semibold mt-6 w-full hover:!bg-gray-400 hover:!text-black"
						onClick={navigateToSignup}
					>
						Sign Up
					</Button>
					<p className="mt-3 text-xs text-gray-500">
						By signing up you agree to the Terms of Service and
						Privacy Policy
					</p>
				</div>
			</div>
		</div>
	);
};

export default Login;
