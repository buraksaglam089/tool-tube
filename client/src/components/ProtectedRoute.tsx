// src/components/ProtectedRoute.tsx
import React from "react";
import { Navigate, Outlet } from "react-router-dom";
import { useGetCurrentUserQuery } from "../app/auth/authApi";

const ProtectedRoute: React.FC = () => {
  const { data: user, isLoading } = useGetCurrentUserQuery({});

  if (isLoading) return <div>Loading...</div>;

  if (!user) {
    // Redirect to login page if not authenticated
    return <Navigate to="/login" replace />;
  }

  return <Outlet />;
};

export default ProtectedRoute;
