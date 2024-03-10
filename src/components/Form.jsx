import React from 'react';
import { Avatar, Button, Form, Input } from 'antd';
import { UserOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';

const ExpenseForm = () => {
	let navigate = useNavigate();

	const onFinish = (values) => {
		console.log('Success:', values);
	};
	const onFinishFailed = (errorInfo) => {
		console.log('Failed:', errorInfo);
	};

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

						<Form.Item label="Category" name="category">
							<Input size="large" />
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
