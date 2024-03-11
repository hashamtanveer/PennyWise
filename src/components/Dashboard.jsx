import React from 'react';
import { useEffect, useState } from 'react';
import { UserOutlined } from '@ant-design/icons';
import { Avatar, Button, Table, Tag } from 'antd';
import { useNavigate } from 'react-router-dom';

const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]

const Dashboard = () => {
	let navigate = useNavigate();
    const [transactions, setTransactions] = useState([]);

    useEffect(() => {
        const token = localStorage.getItem("user_token");
        fetch(`${process.env.REACT_APP_BACKEND_URL}/transactions`, {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${token}`
            }
        }).then(res => {
            if (!res.ok) {
                if (res.status == 401)
                    navigateToLogin();
                return
            }
            res.json().then(data => {
                let balance = 0;
                for (let i = 0; i < data.length; i++) {
                    balance += data[i].amount ? data[i].amount : 0;
                    data[i].key = i;
                    data[i].balance = balance;
                    const date = new Date(data[i].date);
                    data[i].date = `${months[date.getMonth() - 1]} ${date.getDay()}`
                }

                setTransactions(data);
            }).catch(err => {
                console.log("json", err);
            })
        }).catch(err => {
            console.log("fetch", err);
        })
    }, [])

	const navigateToForm = () => {
		navigate('/form');
	  };

    const navigateToLogin = () => {
        navigate('/login');
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

				<div className="text-2xl font-semibold mt-6 flex flex-row justify-between">
					<h2>Recent Transactions</h2>
					<Button onClick={navigateToForm} className=''>Add New Expense</Button>
				</div>

				<div className="mt-6">
					<Table columns={columns} dataSource={transactions} />
				</div>
			</div>
		</div>
	);
};

export default Dashboard;
