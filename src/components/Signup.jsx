import { Button, Form, Input } from 'antd';
import React from 'react';
import { useNavigate } from 'react-router-dom';

const Signup = () => {
    let navigate = useNavigate();

    const onFinish = (values) => {
        fetch(`${process.env.REACT_APP_BACKEND_URL}/user/signup`, {
            headers: {
                'Content-Type': "application/json"
            },
            method: "POST",
            body: JSON.stringify({ username: values.email, password: values.password })
        }).then(res => {
            if (res.status === 200)
                navigateToLogin();
            console.log(res.json());
        }). catch(err => console.log(err))
    };
    const onFinishFailed = (errorInfo) => {
        console.log('Failed:', errorInfo);
    };
    const navigateToLogin = () => {
        navigate('/login');
    };
    return (
        <div className="h-screen">
            {/* Header */}
            <div className="w-full h-16 border-2 border-b-gray-300 flex justify-between items-center px-5">
                <p className="font-semibold text-xl">PennyWise</p>
                <Button type="primary" className="bg-sky-500 font-semibold" onClick={navigateToLogin}>
                    Login
                </Button>
            </div>

            {/* Body */}
            <div className="w-[100%] h-full flex justify-center items-center">
                <div className="w-[500px] h-auto border border-gray-300 p-3 rounded-xl">
                    <p className="text-2xl font-semibold">Sign Up</p>

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
                            <Input size="large" type="password"/>
                        </Form.Item>

                        <Form.Item className="w-full flex justify-end">
                            <Button
                                type="primary"
                                htmlType="submit"
                                className="bg-sky-500"
                                size="large"
                            >
                                Sign Up
                            </Button>
                        </Form.Item>
                    </Form>
                </div>
            </div>
        </div>
    );
};

export default Signup;
