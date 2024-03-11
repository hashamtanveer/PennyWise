import React from 'react';
import { useEffect } from 'react';
import { Avatar, Button, DatePicker, Form, Input } from 'antd';
import { UserOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';

const ExpenseForm = () => {
	let navigate = useNavigate();

    // If user not logged in, kick them back to login
    useEffect(() => {
        const token = localStorage.getItem("user_token");
        fetch(`${process.env.REACT_APP_BACKEND_URL}/transactions`, {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${token}`
            }
        }).then(res => {
            if (!res.ok && res.status === 401)
                navigateToLogin();
        })
    }, [])

	const onFinish = (values) => {
        const token = localStorage.getItem("user_token");
        fetch(`${process.env.REACT_APP_BACKEND_URL}/transactions`, {
            method: "POST",
            headers: {
                'Content-Type': "application/json",
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify({ amount: parseInt(values.amount), category: values.category, description: values.description, date: (values.date === undefined ? new Date(Date.now()) : new Date(values.date)).toISOString() })
        });
    };
	const onFinishFailed = (errorInfo) => {
		console.log('Failed:', errorInfo);
	};

    const navigateToLogin = () => {
        navigate("/login");
    }

	const navigateToDashboard = () => {
		navigate('/dashboard');
	  };

	return (
		<div className="h-screen flex flex-col justify-center items-center">
			{/* Header */}
			<div className="w-full h-16 border-2 border-b-gray-300 flex justify-between items-center px-5">
				<p className="font-semibold text-xl">PennyWise</p>
				<div className="w-[600px] h-full flex justify-between items-center">
					<div className="w-[400px] h-full flex justify-around items-center text-sm">
						<p onClick={navigateToDashboard}>Dashboard</p>
						<p>Invoices</p>
						<p>Reports</p>
						<p>Budgets</p>
					</div>
					<Avatar size="large" icon={<UserOutlined />} />
				</div>
			</div>

			{/* Body */}
			<div className="w-[70%] h-full flex justify-center items-center">
				<div className="w-[500px] h-auto border border-gray-300 p-3 rounded-xl">
					<p className="text-2xl font-semibold">Add An Expense</p>

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
						<Form.Item label="Amount" name="amount">
							<Input size="large" />
						</Form.Item>

						<Form.Item label="Category" name="category">
							<Input size="large" />
						</Form.Item>

						<Form.Item label="Description" name="description">
							<Input size="large" />
						</Form.Item>

						<Form.Item label="Date" name="date">
							<DatePicker />
						</Form.Item>

						<Form.Item className="w-full flex justify-end">
							<Button
								type="primary"
								htmlType="submit"
								className="bg-sky-500"
								size="large"
							>
								Save
							</Button>
						</Form.Item>
					</Form>
				</div>
			</div>
		</div>
	);
};

export default ExpenseForm;
