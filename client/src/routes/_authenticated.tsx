import { createFileRoute, Outlet, redirect } from '@tanstack/react-router';

export const Route = createFileRoute('/_authenticated')({
  beforeLoad(ctx) {
    if (!ctx.context.auth.isAuthenticated) {
      throw redirect({
        to: '/login',
        search: {
          redirect: location.href,
        },
      });
    }
  },
  component() {
    return (
      <div className="p-4">
        <div>common layout</div>
        <Outlet />
      </div>
    );
  },
});
