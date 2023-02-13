class UsersController < ApplicationController
  
  def new
    @flavor = params[:flavor]
    @user = User.new
  end

  def create
    if params[:login] == '1'
      redirect_to dashboard_path, notice: "User was logged in."
      return
    end

    @user = User.new(user_params)
    @user.name = ''
    @flavor = params[:flavor]

    if @user.save
      redirect_to dashboard_path, notice: "User was successfully created."
    else
      render :new, status: :unprocessable_entity
    end
  end

  private

  def user_params
    params.require(:user).permit(:email)
  end
end
