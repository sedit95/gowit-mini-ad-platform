interface ErrorMessageProps {
  message: string;
}

export function ErrorMessage({ message }: ErrorMessageProps) {
  return (
    <div className="error-message" role="alert" aria-live="assertive">
      <p>{message}</p>
    </div>
  );
}
