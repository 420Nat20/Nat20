import Navbar from "./navbar";

export default function Layout({ isAdmin, children }) {
  return (
    <div className="min-h-screen bg-gray-100">
      <Navbar isAdmin={isAdmin} />
      <main className="-mt-24 pb-8">{children}</main>
      <footer>
        <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 lg:max-w-7xl">
          <div className="border-t border-gray-200 py-8 text-sm text-gray-500 text-center sm:text-left">
            <span className="block sm:inline">&copy; 2021 Nat20</span>{" "}
            <span className="block sm:inline">All rights reserved.</span>
          </div>
        </div>
      </footer>
    </div>
  );
}
