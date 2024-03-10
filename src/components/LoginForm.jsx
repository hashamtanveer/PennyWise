import { Button, Form, Input } from 'antd';
import React from 'react';
import { useNavigate } from 'react-router-dom';


const LoginForm = () => {
	let navigate = useNavigate();

	const onFinish = (values) => {
		console.log('Success:', values);
	};
	const onFinishFailed = (errorInfo) => {
		console.log('Failed:', errorInfo);
	};
    const navigateToSignup = () => {
		navigate('/signup');
	  };
      const navigateToDashboard = () => {
		navigate('/dashboard');
	  };
	return (
		<div className="h-screen">
			{/* Header */}
			<div className="w-full h-16 border-2 border-b-gray-300 flex justify-between items-center px-5">
				<p className="font-semibold text-xl">PennyWise</p>
				<Button type="primary" className="bg-sky-500 font-semibold" onClick={navigateToSignup}>
					Sign Up
				</Button>
			</div>

			{/* Body */}
			<div className="w-[100%] h-full flex justify-center items-center">
				<div className="w-[500px] h-auto border border-gray-300 p-3 rounded-xl">
					<p className="text-2xl font-semibold">Login</p>

					<Form
						name="basic"
						layout="vertical"
						initialValues={{
							remember: true,
						}}
						onFinish={onFinish}
						onFinishFailed={onFinishFailed}
						autoComplete="off"
						className="mt-6"
					>
						<Form.Item label="Email" name="email">
							<Input size="large" />
						</Form.Item>

						<Form.Item label="Password" name="password">
							<Input size="large" />
						</Form.Item>

						<Form.Item className="w-full flex justify-end">
							<Button
								type="primary"
								htmlType="submit"
								className="bg-sky-500"
								size="large"
                                onClick={navigateToDashboard}
							>
								Login
							</Button>
						</Form.Item>
					</Form>
				</div>
			</div>
		</div>
	);
};

export default LoginForm;
