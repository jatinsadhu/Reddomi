import Link from "next/link";
import { routes } from "@doota/ui-core/routing";

export default function Page() {
  return (
    <main className="min-h-[70vh] flex items-center justify-center px-6 py-12">
      <div className="w-full max-w-3xl rounded-3xl border border-secondary/30 bg-background/95 p-10 shadow-lg">
        <div className="space-y-4 text-center">
          <p className="text-sm uppercase tracking-[0.2em] text-primary">Post Editor</p>
          <h1 className="text-3xl font-semibold">Editing posts is disabled</h1>
          <p className="text-muted-foreground">
            The post editor is no longer available in the portal UI.
          </p>
          <div className="flex flex-col sm:flex-row justify-center gap-3 pt-4">
            <Link href={routes.new.dashboard} className="inline-flex items-center justify-center rounded-full border border-primary bg-primary/10 px-5 py-3 text-sm font-semibold text-primary transition hover:bg-primary/20">
              Back to dashboard
            </Link>
            <Link href={routes.new.billing} className="inline-flex items-center justify-center rounded-full border border-secondary bg-secondary/10 px-5 py-3 text-sm font-semibold text-secondary transition hover:bg-secondary/20">
              View billing
            </Link>
          </div>
        </div>
      </div>
    </main>
  );
}
