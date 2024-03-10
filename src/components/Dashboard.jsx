import React from 'react';
import { UserOutlined } from '@ant-design/icons';
import { Avatar, Button, Table, Tag } from 'antd';
import { useNavigate } from 'react-router-dom';

const Dashboard = () => {
	let navigate = useNavigate();

	const navigateToForm = () => {
		navigate('/form');
	  };
	const columns = [
		{
			title: 'Date',
			dataIndex: 'date',
			key: 'date',
		},
		{
			title: 'Description',
			dataIndex: 'description',
			key: 'description',
		},
		{
			title: 'Category',
			dataIndex: 'category',
			key: 'category',
			render: (category) => (
				<Tag color="gray" key={category}>
					{category}
				</Tag>
			),
		},
		{
			title: 'Amount',
			key: 'amount',
			dataIndex: 'amount',
		},
		{
			title: 'Balance',
			key: 'balance',
			dataIndex: 'balance',
		},
	];

	const data = [
		{
			key: '1',
			date: 'Jan 1',
			description: 'Monthly Rent',
			category: 'Rent',
			amount: '$1000.00',
			balance: '$12,000.00',
		},
		{
			key: '2',
			date: 'Jan 1',
			description: 'Monthly Rent',
			category: 'Rent',
			amount: '$1000.00',
			balance: '$12,000.00',
		},
		{
			key: '3',
			date: 'Jan 1',
			description: 'Monthly Rent',
			category: 'Rent',
			amount: '$1000.00',
			balance: '$12,000.00',
		},
	];

	return (
		<div className="h-screen flex flex-col justify-center items-center">
			{/* Header */}
			<div className="w-full h-16 border-2 border-b-gray-300 flex justify-between items-center px-5">
				<p className="font-semibold text-xl">PennyWise</p>
				<div className="w-[600px] h-full flex justify-between items-center">
					<div className="w-[400px] h-full flex justify-around items-center text-sm">
						<p>Home</p>
						<p>Reports</p>
						<p>Events</p>
						<p>People</p>
					</div>
					<Avatar size="large" icon={<UserOutlined />} />
				</div>
			</div>

			{/* Body */}
			<div className="w-[70%] h-full">
				<div className="text-3xl font-semibold w-full border-b-2 border-gray-300 p-2">
					Dashboard
				</div>

				<div className="text-2xl font-semibold mt-6">
					Recent Transactions
					<Button onClick={navigateToForm} className='ml-[800px]'>Add New Expense</Button>
				</div>

				<div className="mt-6">
					<Table columns={columns} dataSource={data} />
				</div>
			</div>
		</div>
	);
};

export default Dashboard;
