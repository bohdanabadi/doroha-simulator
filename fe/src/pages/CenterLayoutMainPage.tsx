import React from 'react';

const CenteredLayout = () => {
    return (
        <div className="min-h-screen flex justify-center items-center">
            <div className="w-85">
                {/* Your content goes here */}
                <h1 className="text-3xl mb-4">Welcome to Centered Layout</h1>
                <p>This is the main content of your web application.</p>
            </div>
        </div>
    );
};

export default CenteredLayout;