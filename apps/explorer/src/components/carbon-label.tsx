import { ReactNode } from "react";

export const CarbonLabel = ({
  icon,
  label,
  className,
}: {
  className: string;
  icon: ReactNode;
  label: string;
}) => {
  return (
    <div className={`flex items-center px-5 py-[15px] ${className}`}>
      {icon}

      <span className="ml-[15px] text-[16px] font-medium text-[#fff]">
        {label}
      </span>
    </div>
  );
};
